using dotnetservice.DataAccess;
using dotnetservice.DataAccess.Models;
using dotnetservice.Interfaces.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace dotnetservice.Repositories
{
    public class UserRepository(IServiceProvider serviceProvider) : IUserRepository
    {
        private readonly AppDbContext _context = serviceProvider.GetRequiredService<AppDbContext>();

        public async Task<User?> GetUserByIdAsync(Guid publicId, CancellationToken ctx)
        {
            try
            {
                User? user = await _context.Users.FirstOrDefaultAsync(u => u.PublicId == publicId, ctx);

                return user;
            }
            catch
            {
                return null;
            }
        }

        public async Task<Guid> CreateUserAsync(User user, CancellationToken ctx)
        {
            await _context.Users.AddAsync(user, ctx);
            await _context.SaveChangesAsync(ctx);

            return user.PublicId;
        }

        public async Task<User?> GetUserByEmailAsync(string email, CancellationToken ctx)
        {
            try
            {
                User? user = await _context.Users.FirstOrDefaultAsync(u => u.Email == email, ctx);

                return user;
            }
            catch
            {
                return null;
            }
        }
    }
}