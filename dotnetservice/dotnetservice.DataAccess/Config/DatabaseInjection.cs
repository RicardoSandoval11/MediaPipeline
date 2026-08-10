using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace dotnetservice.DataAccess.Config
{
    public static class DatabaseInjection
    {
        public static IServiceCollection AddDatabase(this IServiceCollection services, string connectionString)
        {
            services.AddDbContext<AppDbContext>(options => options.UseNpgsql(connectionString), ServiceLifetime.Scoped);
            return services;
        }
    }
}
