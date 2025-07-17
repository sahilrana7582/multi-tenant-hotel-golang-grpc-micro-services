import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import type { LoginFormValues } from '../type';
import { login } from '../api';
import toast from 'react-hot-toast';

export const useLogin = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (data: LoginFormValues) => {
    setLoading(true);
    try {
      const { token } = await login(data);

      localStorage.setItem("token", token);

      toast.success("Login successful");

      navigate("/");

    } catch (err: unknown) {
      if (err instanceof Error) {
        toast.error(err.message, {
          duration: 4000,
          position: 'top-center',
        });
      }
    } finally {
      setLoading(false);
    }
  };

  return { handleLogin, loading };
};
