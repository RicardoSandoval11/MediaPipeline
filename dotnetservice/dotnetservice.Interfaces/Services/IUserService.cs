using dotnetservice.Models.Requests;
using dotnetservice.Models.Responses;

namespace dotnetservice.Interfaces.Services
{
    public interface IUserService
    {
        Task<CreateUserResponse> CreateUserAsync(CreateUserRequest request, CancellationToken ctx);

        Task<AuthenticateUserResponse> AuthenticateUserAsync(AuthenticateUserRequest request, CancellationToken ctx);
    }
}