import { useEffect, useState } from "react";
import type { Department } from "../types";
import type { ApiResponse } from "../../../types/types";
import apiClient from "../../../services/apiClient";

export function useDepartments() {
  const [departments, setDepartments] = useState<Department[] | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchDepartments = async () => {
      try {
        const response = await apiClient.get<ApiResponse<Department[]>>("/department/77ec8de4-8251-4a92-9db1-b353d57ce4b1");
        setDepartments(response.data.data);
      } 
      
      catch (err: unknown) {
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError("An unknown error occurred.");
        }
      } finally {
        setLoading(false);
      }
    };

    fetchDepartments();
  }, []);

  return { departments, loading, error };
}
