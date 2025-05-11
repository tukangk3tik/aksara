export interface RequestHeaders {
  'Content-Type': string;
  Authorization?: string;
  [key: string]: string | undefined;
}