import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:3000",
  withCredentials: true, // Ensures cookies (refresh token) are sent
});

// Request Interceptor: Attach access token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("access_token");
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});

// Response Interceptor: Handle expired tokens
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      try {
        // Call refresh token API
        const res = await axios.post("http://localhost:3000/refresh", {}, { withCredentials: true });
        localStorage.setItem("AccessToken", res.data.access_token);

        // Retry the failed request with new access token
        error.config.headers["Authorization"] = `Bearer ${res.data.access_token}`;
        return api.request(error.config);
      } catch (refreshError) {
        console.error("Session expired. Please log in again.");
        localStorage.removeItem("access_token");
        localStorage.removeItem("username");
        localStorage.removeItem("job_title");
        window.location.href = "/login"; // Redirect to login page
      }
    }
    return Promise.reject(error);
  }
);

export default api;