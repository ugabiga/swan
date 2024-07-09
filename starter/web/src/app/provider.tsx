'use client';

import React from "react";
import {ThemeProvider} from "@/lib/theme-provider";
import {QueryClientProvider} from "@tanstack/react-query";
import queryClient from "@/lib/query-client";
import {Provider} from "react-redux";
import {store} from "@/store";

export default function Providers({children}: { children: React.ReactNode }) {
    return (
        <ThemeProvider attribute="class" defaultTheme={"dark"} enableSystem>
            <QueryClientProvider client={queryClient}>
                <Provider store={store}>
                    {children}
                </Provider>
            </QueryClientProvider>
        </ThemeProvider>
    )
}