import axios from "axios";

let GetProfileByUserID = function(token?: string, userid?: string) {
  return new Promise(function(resolve, reject) {
    if (!token) {
      return reject("Unauthorized");
    }
    axios
      .get(`/api/profile/${userid}`, { 
        headers: { Authorization: token }
      })
      .then((res) => {
        return resolve(res.data);
      })
      .catch((err) => {
        if (err.response.status === 401) {
          return reject("Unauthorized");
        } else if (err.response.data.msg) {
          return reject(err.response.data.msg);
        }
        return reject("Unknown error");
      });
  });
};

export default GetProfileByUserID;