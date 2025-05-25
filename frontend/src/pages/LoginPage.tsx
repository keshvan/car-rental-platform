import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../components/AuthContext';
import { authService } from '../services/authService';


const isValidEmail = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g;

export default function LoginPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const { setUser } = useAuth();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setIsLoading(true);
    setError(null);

    if (email && email.match(isValidEmail) && password) {
      try {
        const res = await authService.login({ email, password });
        setUser(res.user);
        navigate("/");
      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
        }
      }
    } else {
      setError('Введите корректный email и пароль');
    }

    setIsLoading(false);
  };

  return (
    <div className="flex flex-col items-center justify-start pt-12 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8 p-8 bg-white shadow-lg rounded-md">
        <div>
          <h2 className="text-center text-2xl font-bold">
            Войти
          </h2>
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          <div>
            <label htmlFor="username" className="block text-sm font-medium mb-1">
              Email
            </label>
            <input
              id="username"
              name="username"
              type="text"
              autoComplete="username"
              required
              className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium mb-1">
              Пароль
            </label>
            <input
              id="password"
              name="password"
              type="password"
              autoComplete="new-password"
              required
              className="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          {error && (
            <p className="text-sm text-red-600 text-center">{error}</p>
          )}

          <div>
            <button
              type="submit"
              disabled={isLoading}
              className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 disabled:bg-blue-300"
            >
              {isLoading ? 'Вход...' : 'Войти'}
            </button>
          </div>
          <div className="text-sm text-center">
            <Link to="/register" className="font-medium text-blue-600 hover:text-blue-500 hover:underline">
              Нет аккаунта? Зарегестрироваться
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
};