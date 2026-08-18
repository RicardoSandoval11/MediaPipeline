using System.Text.Json.Serialization;

namespace dotnetservice.Models.Responses
{
    public class LivenessResponse
    {
        [JsonPropertyName("message")]
        public string Message { get; set; } = null!;
    }
}