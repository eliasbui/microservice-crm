using AuthService.Models;
using AuthService.Models.DTOs;

namespace AuthService.Services;

public interface IUserService
{
    Task<LoginResponse?> LoginAsync(LoginRequest request);
    Task<UserDto?> RegisterAsync(RegisterRequest request);
    Task<LoginResponse?> RefreshTokenAsync(RefreshTokenRequest request);
    Task<bool> RevokeTokenAsync(string userId);
}
