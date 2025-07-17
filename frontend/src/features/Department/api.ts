import apiClient from "../../services/apiClient";
import type { Department } from "./types";
import type { ApiResponse } from "../../types/types";

export const getDepartments = async (): Promise<Department[]> => {
  const response = await apiClient.get<ApiResponse<Department[]>>('/departments');
  return response.data.data;
};