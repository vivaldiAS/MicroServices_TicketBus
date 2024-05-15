<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class RoleMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @param  string  ...$role
     * @return mixed
     */
    public function handle(Request $request, Closure $next, ...$role)
    {
        $user = Auth::user();

        if (!$user) {
            return response()->json(['message' => 'Unauthorized'], 401);
        }

        if (!in_array($user->role->role, $role)) {
            return response()->json(['message' => 'You do not have permission to access this page.'], 403);
        }

        return $next($request);
    }
}
