using dotnetservice.DataAccess.Models;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace dotnetservice.DataAccess.Config
{
    public class FileCounterConfig : IEntityTypeConfiguration<FileCounter>
    {
        public void Configure(EntityTypeBuilder<FileCounter> builder)
        {
            builder.ToTable("file_counter");

            builder.HasKey(e => e.Id);

            builder.Property(e => e.Id)
            .HasColumnName("id")
            .ValueGeneratedNever()
            .IsRequired();

            builder.Property(e => e.Count)
            .HasColumnName("count")
            .IsRequired();

            builder.HasOne(e => e.User)
            .WithMany(u => u.FileCounters)
            .HasForeignKey(e => e.UserId)
            .IsRequired();

            builder.Property(e => e.UserId)
            .HasColumnName("user_id")
            .IsRequired();

            builder.Property(e => e.StartDate)
            .HasColumnName("start_date")
            .IsRequired();

            builder.Property(e => e.EndDate)
            .HasColumnName("end_date")
            .IsRequired();
        }
    }
}