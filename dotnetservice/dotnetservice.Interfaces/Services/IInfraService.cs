using dotnetservice.Models.Responses;

namespace dotnetservice.Interfaces.Services
{
    public interface IInfraService
    {
        LivenessResponse Liveness();

        ReadinessResponse Readiness();
    }
}