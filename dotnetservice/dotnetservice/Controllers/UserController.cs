
namespace dotnetservice.Controllers
{
    using System.Net;
    using dotnetservice.Interfaces.Services;
    using dotnetservice.Models.Requests;
    using dotnetservice.Models.Responses;
    using Microsoft.AspNetCore.Mvc;

    [Route("api/v1/user")]
    [ApiController]
    public class UserController(IServiceProvider serviceProvider) : ControllerBase
    {
        private readonly IUserService _service = serviceProvider.GetRequiredService<IUserService>();

        [HttpPost("Create")]
        [Consumes("application/json")]
        public async Task<IActionResult> CreateUserAsync([FromBody] CreateUserRequest request, CancellationToken ctx = default)
        {
            try
            {
                var response = await _service.CreateUserAsync(request, ctx);

                return StatusCode((int)HttpStatusCode.OK, response);
            }
            catch
            {
                return StatusCode((int)HttpStatusCode.InternalServerError, new CreateUserResponse()
                {
                    Success = false,
                    Error = "Something went wrong"
                });
            }
        }

        [HttpPost("Authenticate")]
        public async Task<IActionResult> AuthenticateUserAsync([FromBody] AuthenticateUserRequest request, CancellationToken ctx = default)
        {
            try
            {
                var response = await _service.AuthenticateUserAsync(request, ctx);

                return StatusCode((int)HttpStatusCode.OK, response);
            }
            catch
            {
                return StatusCode((int)HttpStatusCode.InternalServerError, new AuthenticateUserResponse()
                {
                    Success = false,
                    Error = "Something went wrong",
                });
            }
        }
    }
}