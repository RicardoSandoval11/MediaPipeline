using dotnetservice.DataAccess.Models;

namespace dotnetservice.Interfaces.Repositories
{
    public interface IFileCounterRepository
    {
        Task<FileCounter> UpsertFileCounterAsync(FileCounter fileCounter, CancellationToken ctx);

        Task<FileCounter?> GetFileCounterAsync(long userId, DateTime time, CancellationToken ctx);
    }
}