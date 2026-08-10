using dotnetservice.DataAccess.Config;
using dotnetservice.DataAccess.Models;
using Microsoft.EntityFrameworkCore;

namespace dotnetservice.DataAccess
{
    public class AppDbContext(DbContextOptions options) : DbContext(options)
    {
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.ApplyConfiguration(new UserConfig());
            modelBuilder.ApplyConfiguration(new FileCounterConfig());
            base.OnModelCreating(modelBuilder);
        }

        public virtual DbSet<User> Users { get; set; }

        public virtual DbSet<FileCounter> FileCounters { get; set; }
    }
}