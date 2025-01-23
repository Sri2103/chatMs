import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
interface ApiRequestOptions extends AxiosRequestConfig {
	// Add any custom properties you need for your requests here
	// For example:
	// headers?: {
	//   Authorization?: string;
	// }
}

interface ApiResponse<T> {
	data: T;
	status: number;
	statusText: string;
}

class Api {
	private instance: AxiosInstance;

	constructor(baseURL: string) {
		this.instance = axios.create({
			baseURL,
		});
	}

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	private async request<T, D = any>(
		method: 'get' | 'post' | 'put' | 'delete' | 'patch', // Added 'patch'
		endpoint: string,
		data?: D,
		config?: ApiRequestOptions
	): Promise<ApiResponse<T>> {
		try {
			const response = await this.instance[method]<T>(endpoint, data, config);
			return this.formatResponse(response);
		} catch (error) {
			console.error('API Error:', error);
			throw error;
		}
	}

	get<T>(endpoint: string, config?: ApiRequestOptions): Promise<ApiResponse<T>> {
		return this.request<T>('get', endpoint, undefined, config);
	}

	post<T, D>(endpoint: string, data?: D, config?: ApiRequestOptions): Promise<ApiResponse<T>> {
		return this.request<T, D>('post', endpoint, data, config);
	}

	put<T, D>(endpoint: string, data?: D, config?: ApiRequestOptions): Promise<ApiResponse<T>> {
		return this.request<T, D>('put', endpoint, data, config);
	}

	delete<T>(endpoint: string, config?: ApiRequestOptions): Promise<ApiResponse<T>> {
		return this.request<T>('delete', endpoint, undefined, config);
	}

	patch<T, D>(endpoint: string, data?: D, config?: ApiRequestOptions): Promise<ApiResponse<T>> {
		return this.request<T, D>('patch', endpoint, data, config);
	}

	private formatResponse<T>(response: AxiosResponse<T>): ApiResponse<T> {
		return {
			data: response.data,
			status: response.status,
			statusText: response.statusText,
		};
	}
}

export default Api;
