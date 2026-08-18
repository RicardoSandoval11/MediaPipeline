using System.Net;
using dotnetservice.Interfaces.Services;
using dotnetservice.Models.Responses;
using Microsoft.AspNetCore.Mvc;

namespace dotnetservice.Controllers
{
    [Route("api/v1/infra")]
    [ApiController]
    public class InfraController(IServiceProvider serviceProvider) : ControllerBase
    {
        private readonly IInfraService _service = serviceProvider.GetRequiredService<IInfraService>();

        [HttpGet("/livez")]
        [Produces("application/json")]
        public IActionResult Liveness()
        {
            LivenessResponse response = _service.Liveness();

            return StatusCode((int)HttpStatusCode.OK, response);
        }

        [HttpGet("/readyz")]
        [Produces("application/json")]
        public IActionResult Readiness()
        {
            ReadinessResponse response = _service.Readiness();

            return StatusCode((int)HttpStatusCode.OK, response);
        }
    }
}