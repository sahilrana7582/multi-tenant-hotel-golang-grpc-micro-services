import axios from "axios";


const apiClient = axios.create({
    baseURL: "http://localhost:8080/api",
    withCredentials: true,
    timeout: 10000,
    headers: {
      'Content-Type': 'application/json',
    },
});


apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
      if (!error.response) {
        return Promise.reject(new Error("Network error. Please check your connection."));
      }
  
      const res = error.response;
      const errorMessage =
        res.data?.message || "An unknown error occurred. Please try again.";
  
      return Promise.reject(new Error(errorMessage));
    }
  );

export default apiClient;
