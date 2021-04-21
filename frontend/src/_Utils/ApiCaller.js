const HEADERS = {
  "Content-Type": "application/json",
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, PUT, POST, DELETE, HEAD, OPTIONS",
};

const METHODS = {
  GET: "GET",
  POST: "POST",
  DELETE: "DELETE",
};

const STATUS_CODES = {
  ISE: 503,
  OK: 200,
};

const MODE = "cors";

function fetchAndResolve(link, opt) {
  return fetch(link, opt)
    .then((response) => response.json())
    .then((data) => {
      if (data.status !== STATUS_CODES.OK && data.status !== undefined) {
        throw new Error(data.msg);
      }

      return data;
    })
    .catch((error) => {
      console.log(opt.method, link, error);
    });
}

export function Get(link) {
  const opt = {
    method: METHODS.GET,
    headers: HEADERS,
    mode: MODE,
  };

  return fetchAndResolve(link, opt);
}

export function Post(link, body) {
  const opt = {
    method: METHODS.POST,
    headers: HEADERS,
    mode: MODE,
    body: JSON.stringify(body),
  };

  return fetchAndResolve(link, opt);
}

export function Delete(link) {
  const opt = {
    method: METHODS.DELETE,
    headers: HEADERS,
    mode: MODE,
  };

  return fetchAndResolve(link, opt);
}
