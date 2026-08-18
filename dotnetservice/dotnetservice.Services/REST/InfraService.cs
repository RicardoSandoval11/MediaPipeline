using dotnetservice.Interfaces.Services;
using dotnetservice.Models.Responses;

namespace dotnetservice.Services.REST
{
    public class InfraService : IInfraService
    {
        public LivenessResponse Liveness()
        {
            return new()
            {
                Message = "ok"
            };
        }

        public ReadinessResponse Readiness()
        {
            return new()
            {
                Message = "ok"
            };
        }
    }
}