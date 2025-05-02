import axios from "axios";
import Cookies from "js-cookie";

let GetProfile = new Promise(function(resolve, reject) {
  const token = Cookies.get("token");
  if (!token) {
    return reject("Unauthorized");
  }
  axios
    .get("/api/profile", { headers: { Authorization: token } })
    .then((res) => {
      return resolve(res.data);
    })
    .catch((err) => {
      if (err.response.status === 401) {
        return reject("Unauthorized");
      }
      return reject("Unknown error");
    });
    
})

export default GetProfile;