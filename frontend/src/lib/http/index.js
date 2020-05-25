export class FetchError extends Error {
  constructor(message, response) {
    super(message);
    this.name = this.constructor.name;
    this.response = response;
  }
}

export const fetchX = async (url, options) => {
  const response = await fetch(url, options);

  if (!response.ok) {
    throw new FetchError(null, response);
  }

  const contentType = response.headers.get('Content-Type');

  if (contentType.includes('application/json')) {
    try {
      const json = await response.json();
      return json;
    } catch (error) {
      console.error(error);
      return;
    }
  }

  const text = await response.text();
  return text;
};
