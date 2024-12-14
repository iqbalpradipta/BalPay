export interface UserState {
    id: number | null;
    name: string | null;
    email: string | null;
    password: string | null;
    phoneNumber: string | null;
    role: string | null;
    status: 'idle' | 'loading' | 'success' | 'error';
}