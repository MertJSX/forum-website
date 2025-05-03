import axios from "axios";

let GetProfile = function(token?: string) {
  return new Promise(function(resolve, reject) {
    if (!token) {
      return reject("Unauthorized");
    }
    axios
      .get("/api/profile", { 
        headers: { Authorization: token }
      })
      .then((res) => {
        return resolve(res.data);
      })
      .catch((err) => {
        if (err.response.status === 401) {
          return reject("Unauthorized");
        }else if (err.response.msg) {
          return reject(err.response.msg);
        }
        return reject("Unknown error");
      });
  });
};

export default GetProfile;