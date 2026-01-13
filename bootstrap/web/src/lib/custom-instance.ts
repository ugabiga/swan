import Axios, {AxiosError, AxiosRequestConfig} from 'axios';

export const AXIOS_INSTANCE = Axios.create({}); // use your own URL here or environment variable

// add a second `options` argument here if you want to pass extra options to each generated query
export const customInstance = <T>(
    config: AxiosRequestConfig,
    options?: AxiosRequestConfig,
): Promise<T> => {
    const controller = new AbortController();
    const promise = AXIOS_INSTANCE({
        ...config,
        ...options,
        signal: controller.signal,
        withCredentials: true,
    }).then(({data}) => data);

    // @ts-expect-error - Adding cancel method for compatibility
    promise.cancel = () => {
        controller.abort();
    };

    return promise;
};

export type ErrorType<Error> = AxiosError<Error>;
