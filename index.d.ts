declare module 'github-to-cdn' {
  export default function ghCDN(link: string | Partial<{
    username: string;
    repo: string;
    path: string;
    branch: string;
  }>): string;
}
