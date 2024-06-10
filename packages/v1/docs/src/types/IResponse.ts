interface IResponse {
  body: string;
  headers: Record<string, string>;
  statusCode: number;
}

export default IResponse;
