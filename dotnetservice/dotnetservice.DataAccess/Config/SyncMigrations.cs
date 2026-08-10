using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace dotnetservice.DataAccess.Config
{
    public class SyncMigrations(IServiceProvider serviceProvider)
    {
        private readonly IServiceProvider _serviceProvider = serviceProvider;

        public void SyncPendingMigrations()
        {
            try
            {
                using var scope = this._serviceProvider.CreateScope();
                using AppDbContext ctx = scope.ServiceProvider.GetRequiredService<AppDbContext>();
                ctx.Database.Migrate();
            }
            catch
            {
                throw;
            }
        }
    }
}