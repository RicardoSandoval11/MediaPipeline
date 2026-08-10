using dotnetservice.Models.Requests;
using FluentValidation;

namespace dotnetservice.Models.Validators
{
    public class CreateUserValidator : AbstractValidator<CreateUserRequest>
    {
        public CreateUserValidator()
        {
            RuleFor(r => r.Email)
            .Cascade(CascadeMode.Stop)
            .NotNull()
            .WithMessage("Email address is required")
            .NotEmpty()
            .WithMessage("Email address can't be empty")
            .EmailAddress(FluentValidation.Validators.EmailValidationMode.AspNetCoreCompatible)
            .WithMessage("Email address is not valid");

            RuleFor(r => r.Password)
            .Cascade(CascadeMode.Stop)
            .NotNull()
            .WithMessage("Password is required")
            .NotEmpty()
            .WithMessage("Password can't be empty");
        }
    }
}