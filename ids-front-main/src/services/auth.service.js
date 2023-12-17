import axios from "axios";

const API_URL = "http://localhost:8080/";

class AuthService {
  login(email, password) {
    return axios
      .post(API_URL + "login", {
        email,
        password,
      })
      .then((response) => {
        if (response.data.accessToken) {
          localStorage.setItem("access_token", response.data.accessToken);
        }

        return response.data;
      });
  }

  logout() {
    localStorage.removeItem("access_token");
  }

  register(username, email, password) {
    return axios
      .post(API_URL + "signup", {
        name: username,
        email,
        password,
      })
      .then((response) => {
        if (response.data.accessToken) {
          localStorage.setItem("access_token", response.data.accessToken);
        }

        return response.data;
      });
  }

  getAccessToken() {
    return localStorage.getItem("access_token");
  }
}

export default new AuthService();
