import ReactDOM from 'react-dom/client'
import {createRouter, RouterProvider} from '@tanstack/react-router'
import {routeTree} from './routeTree.gen'
import {ThemeProvider} from "@/components/theme-provider";
import {QueryClientProvider} from "@tanstack/react-query";
import queryClient from "@/lib/query-client";
import { Toaster } from "@/components/ui/sonner";

// Set up a Router instance
const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
})

// Register things for typesafety
declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

const rootElement = document.getElementById('app')!

if (!rootElement.innerHTML) {
    const root = ReactDOM.createRoot(rootElement)
    root.render(
        <QueryClientProvider client={queryClient}>
            <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
                <RouterProvider router={router}/>
                <Toaster />
            </ThemeProvider>
        </QueryClientProvider>
    )
}