using System.Text.Json.Serialization;

namespace dotnetservice.Models.Responses
{
    public class AuthenticateUserResponse
    {
        [JsonPropertyName("success")]
        public bool Success { get; set; }

        [JsonPropertyName("error")]
        public string? Error { get; set; }

        [JsonPropertyName("token")]
        public string? Token { get; set; }
    }
}