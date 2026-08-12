namespace dotnetservice.DataAccess.Models
{
    public class User
    {
        public long Id { get; set; }

        public string Email { get; set; } = null!;

        public string Password { get; set; } = null!;

        public Guid PublicId { get; set; }

        public virtual ICollection<FileCounter> FileCounters { get; set; } = new List<FileCounter>();
    }
}