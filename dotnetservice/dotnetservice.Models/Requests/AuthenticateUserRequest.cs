using System.Text.Json.Serialization;

namespace dotnetservice.Models.Requests
{
    public class AuthenticateUserRequest
    {
        [JsonPropertyName("email")]
        public string Email { get; set; } = null!;

        [JsonPropertyName("password")]
        public string Password { get; set; } = null!;
    }
}