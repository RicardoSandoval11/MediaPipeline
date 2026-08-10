using System.Text.Json.Serialization;

namespace dotnetservice.Models.Responses
{
    public class CreateUserResponse
    {
        [JsonPropertyName("success")]
        public bool Success { get; set; }

        [JsonPropertyName("user_id")]
        public Guid? UserId { get; set; }

        [JsonPropertyName("error")]
        public string? Error { get; set; }
    }
}