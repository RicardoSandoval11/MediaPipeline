using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Security.Cryptography;
using System.Text;
using dotnetservice.DataAccess.Models;
using dotnetservice.Interfaces.Repositories;
using dotnetservice.Interfaces.Services;
using dotnetservice.Models.Requests;
using dotnetservice.Models.Responses;
using dotnetservice.Models.Validators;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.IdentityModel.Tokens;

namespace dotnetservice.Services.REST
{
    public class UserService(IServiceProvider serviceProvider) : IUserService
    {

        private readonly IUserRepository _repository = serviceProvider.GetRequiredService<IUserRepository>();
        private readonly string _secretKey = "TuClaveSuperSecretaQueDebeSerMuyLargaYSeguraDeAlMenos256Bits";
        private readonly string _issuer = "MiAppApi";
        private readonly string _audience = "MiAppClientes";

        public async Task<AuthenticateUserResponse> AuthenticateUserAsync(AuthenticateUserRequest request, CancellationToken ctx)
        {
            var validator = new AuthenticateUserValidator();

            var validationResult = validator.Validate(request);

            if (validationResult.Errors.Count > 0)
            {
                return new AuthenticateUserResponse()
                {
                    Success = false,
                    Error = validationResult.Errors[0].ErrorMessage,
                };
            }

            User? user = await _repository.GetUserByEmailAsync(request.Email, ctx);

            if (user == null)
            {
                return new AuthenticateUserResponse()
                {
                    Success = false,
                    Error = "Email does not exist",
                };
            }

            string pwdHash = HashString(request.Password);

            if (pwdHash != user.Password)
            {
                return new AuthenticateUserResponse()
                {
                    Success = false,
                    Error = "Invalid password",
                };
            }

            return new AuthenticateUserResponse()
            {
                Success = true,
                Token = GenerateTokenJwt(user.PublicId.ToString()),
            };
        }

        public async Task<CreateUserResponse> CreateUserAsync(CreateUserRequest request, CancellationToken ctx)
        {

            var validator = new CreateUserValidator();

            var validationResult = validator.Validate(request);

            if (validationResult.Errors.Count > 0)
            {
                return new()
                {
                    Success = false,
                    Error = validationResult.Errors[0].ErrorMessage,
                };
            }

            User? existingUser = await _repository.GetUserByEmailAsync(request.Email, ctx);

            if (existingUser != null)
            {
                return new()
                {
                    Success = false,
                    Error = "Email already exists"
                };
            }

            User newUser = new()
            {
                PublicId = Guid.NewGuid(),
                Email = request.Email,
                Password = HashString(request.Password),
            };

            Guid userId = await _repository.CreateUserAsync(newUser, ctx);

            return new()
            {
                Success = true,
                UserId = userId,
            };
        }

        private static string HashString(string input)
        {
            byte[] bytes = Encoding.UTF8.GetBytes(input);

            byte[] hashBytes = SHA256.HashData(bytes);

            return Convert.ToHexString(hashBytes).ToLower();
        }

        private string GenerateTokenJwt(string userId)
        {
            var claims = new[]
            {
                new Claim(ClaimTypes.NameIdentifier, userId),
                new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString())
            };

            var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_secretKey));
            var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

            var tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(claims),
                Expires = DateTime.UtcNow.AddHours(2),
                Issuer = _issuer,
                Audience = _audience,
                SigningCredentials = creds
            };

            var tokenHandler = new JwtSecurityTokenHandler();
            var token = tokenHandler.CreateToken(tokenDescriptor);

            return tokenHandler.WriteToken(token);
        }
    }
}