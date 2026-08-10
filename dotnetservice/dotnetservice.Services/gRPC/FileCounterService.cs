using dotnetservice.DataAccess.Models;
using dotnetservice.Interfaces.Repositories;
using dotnetservice.Services.Protos;
using Grpc.Core;
using Microsoft.Extensions.DependencyInjection;

namespace dotnetservice.Services.gRPC
{
    public class FileCounterService(IServiceProvider serviceProvider) : Protos.FileCounterService.FileCounterServiceBase
    {
        private readonly IFileCounterRepository _fileCounterRepository = serviceProvider.GetRequiredService<IFileCounterRepository>();
        private readonly IUserRepository _userRepository = serviceProvider.GetRequiredService<IUserRepository>();
        private readonly int MAX_FILES_ALLOWED = 5;

        public override async Task<UpsertFileCounterResponse> UpsertFileCounterAsync(UpsertFileCounterRequest request, ServerCallContext ctx)
        {
            if (!Guid.TryParse(request.UserId, out Guid publicId))
            {
                return new UpsertFileCounterResponse()
                {
                    Success = false,
                    Error = "Invalid userId"
                };
            }

            User? user = await _userRepository.GetUserByIdAsync(publicId, ctx.CancellationToken);

            if (user == null)
            {
                return new UpsertFileCounterResponse()
                {
                    Success = false,
                    Error = "User does not exist"
                };
            }

            DateTime time = DateTime.UtcNow;

            FileCounter? current = await _fileCounterRepository.GetFileCounterAsync(user.Id, time, ctx.CancellationToken);

            if (current != null && current.Count >= MAX_FILES_ALLOWED)
            {
                return new UpsertFileCounterResponse()
                {
                    Success = true,
                    LimitReached = true
                };
            }

            if (current != null && current.Count < MAX_FILES_ALLOWED)
            {
                current.Count++;
                await _fileCounterRepository.UpsertFileCounterAsync(current, ctx.CancellationToken);

                return new UpsertFileCounterResponse()
                {
                    Success = true,
                    LimitReached = false
                };
            }

            FileCounter newCounter = new()
            {
                Count = 1,
                StartDate = time,
                EndDate = time.AddHours(24),
                UserId = user.Id,
                Id = Guid.NewGuid()
            };

            await _fileCounterRepository.UpsertFileCounterAsync(newCounter, ctx.CancellationToken);

            return new UpsertFileCounterResponse()
            {
                Success = true,
                LimitReached = false
            };
        }
    }
}