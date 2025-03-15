
export async function authorizedReq(url, options = {}) {
    const token = localStorage.getItem("access_token") || "";
  
    return fetch(url, {
      ...options,
      headers: {
        ...options.headers,
        "Content-Type": "application/json",
        "Accept": "application/json",
        Authorization: `Bearer ${token}`
      }
    });
  }