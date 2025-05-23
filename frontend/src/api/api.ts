import axios from "axios";

export const carApi = axios.create({baseURL: import.meta.env.VITE_CAR_URL})