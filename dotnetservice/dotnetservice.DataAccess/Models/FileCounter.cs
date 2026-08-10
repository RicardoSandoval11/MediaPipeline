namespace dotnetservice.DataAccess.Models
{
    public class FileCounter
    {
        public Guid Id { get; set; }

        public int Count { get; set; }

        public long UserId { get; set; }

        public DateTime StartDate { get; set; }

        public DateTime EndDate { get; set; }

        public virtual User User { get; set; } = null!;
    }
}