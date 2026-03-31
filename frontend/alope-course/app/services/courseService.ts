import api from "../lib/axios";

export const getCourses = async () => {
  try {
    const response = await api.get("/ping");
    return response.data;
  } catch (error) {
    throw error;
  }
};
