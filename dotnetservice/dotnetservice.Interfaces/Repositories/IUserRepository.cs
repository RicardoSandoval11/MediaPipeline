using dotnetservice.DataAccess.Models;

namespace dotnetservice.Interfaces.Repositories
{
    public interface IUserRepository
    {
        Task<User?> GetUserByIdAsync(Guid publicId, CancellationToken ctx);

        Task<Guid> CreateUserAsync(User user, CancellationToken ctx);

        Task<User?> GetUserByEmailAsync(string email, CancellationToken ctx);
    }
}