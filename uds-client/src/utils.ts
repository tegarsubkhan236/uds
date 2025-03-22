export const handleError = (err: unknown): string => {
    if (err instanceof Error) {
        return err.message;
    }
    return 'An unknown error occurred';
}
