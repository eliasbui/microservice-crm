using UserAuthService.Models;

namespace UserAuthService.Services;

public interface IAuthService
{
    Task<(bool Success, string? Error, AuthResponse? Response)> RegisterAsync(RegisterRequest request);
    Task<(bool Success, string? Error, AuthResponse? Response)> LoginAsync(LoginRequest request);
    Task<(bool Success, string? Error, AuthResponse? Response)> RefreshTokenAsync(RefreshTokenRequest request);
    Task<(bool Success, string? Error)> ChangePasswordAsync(string userId, ChangePasswordRequest request);
    Task<bool> RevokeTokenAsync(string userId);
}
