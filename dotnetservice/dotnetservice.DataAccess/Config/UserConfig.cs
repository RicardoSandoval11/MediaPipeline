using dotnetservice.DataAccess.Models;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace dotnetservice.DataAccess.Config
{
    public class UserConfig : IEntityTypeConfiguration<User>
    {
        public void Configure(EntityTypeBuilder<User> builder)
        {
            builder.ToTable("app_user");

            builder.HasKey(e => e.Id);

            builder.Property(e => e.Id)
            .HasColumnName("id")
            .ValueGeneratedOnAdd()
            .IsRequired();

            builder.Property(e => e.Email)
            .HasColumnName("email")
            .IsRequired();

            builder.Property(e => e.Password)
            .HasColumnName("password")
            .IsRequired();

            builder.HasIndex(e => new { e.Password, e.Email });

            builder.HasIndex(e => e.Email).IsUnique();

            builder.Property(e => e.PublicId)
            .HasColumnName("public_id")
            .IsRequired();
        }
    }
}