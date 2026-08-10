using dotnetservice.DataAccess;
using dotnetservice.DataAccess.Models;
using dotnetservice.Interfaces.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace dotnetservice.Repositories
{
    public class FileCounterRepository(IServiceProvider serviceProvider) : IFileCounterRepository
    {
        private readonly AppDbContext _context = serviceProvider.GetRequiredService<AppDbContext>();

        public async Task<FileCounter?> GetFileCounterAsync(long userId, DateTime time, CancellationToken ctx)
        {
            try
            {
                return await _context.FileCounters.FirstOrDefaultAsync(e =>
                    e.UserId == userId &&
                    e.StartDate <= time &&
                    e.EndDate >= time
                , ctx);
            }
            catch
            {
                return null;
            }
        }

        public async Task<FileCounter> UpsertFileCounterAsync(FileCounter fileCounter, CancellationToken ctx)
        {
            await using var transaction = await _context.Database.BeginTransactionAsync(ctx);

            try
            {
                DateTime currentDate = DateTime.UtcNow;
                FileCounter? current = await _context.FileCounters
                .FirstOrDefaultAsync(e => e.UserId == fileCounter.UserId && e.StartDate <= currentDate && e.EndDate >= currentDate, ctx);

                if (current == null)
                {
                    await _context.FileCounters.AddAsync(fileCounter, ctx);
                }

                await _context.SaveChangesAsync(ctx);

                await transaction.CommitAsync(ctx);

                return current ?? fileCounter;
            }
            catch
            {
                await transaction.RollbackAsync(ctx);
                throw;
            }
        }
    }
}