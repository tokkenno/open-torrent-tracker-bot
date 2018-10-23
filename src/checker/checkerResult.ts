export interface CheckerResult {
    online: boolean; // TRUE if the web respond the request
    regOpened: boolean; // TRUE if the verification grants that the registry is open
    regClosed: boolean; // TRUE id the verification grants that the registry is close
    error?: string;
}