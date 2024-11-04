import {QueryClient} from "@tanstack/query-core";
import {AxiosError} from "axios";

const retryStatusCodes = new Set([500]);

const retryLogic = (failureCount: number, error: unknown): boolean => {
    if (!isAxiosError(error)) {
        return false
    }

    // noinspection RedundantIfStatementJS
    if (error.response?.status && retryStatusCodes.has(error.response.status) && failureCount < 3) {
        return true;
    }

    return false;
};

function isAxiosError(error: unknown): error is AxiosError {
    return (error as AxiosError).isAxiosError !== undefined;
}


const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            retry: retryLogic,
        },
    },
});
export default queryClient;
