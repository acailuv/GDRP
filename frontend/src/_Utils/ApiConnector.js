import link from "./backend-link.txt";

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

const MODE = "cors";

const DEFAULT_OPTIONS = {
  headers: HEADERS,
  mode: MODE,
};

const STATUS_CODES = {
  ISE: 503,
  OK: 200,
};

var apiLinkCache = "";

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

export async function GetAPILink() {
  if (apiLinkCache !== "") {
    return apiLinkCache;
  }

  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.GET,
  };

  const backendLink = await fetch(link, opt)
    .then((response) => response.text())
    .then((data) => {
      return data.split(" ")[3].replace("\n", "");
    })
    .catch((err) => console.log(err));

  apiLinkCache = backendLink;

  return backendLink;
}

export function Get(link) {
  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.GET,
  };

  return fetchAndResolve(link, opt);
}

export function Post(link, body) {
  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.POST,
    body: JSON.stringify(body),
  };

  return fetchAndResolve(link, opt);
}

export function Delete(link) {
  const opt = {
    ...DEFAULT_OPTIONS,
    method: METHODS.DELETE,
  };

  return fetchAndResolve(link, opt);
}
