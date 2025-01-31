import { createFileRoute } from "@tanstack/react-router";
import "../index.css";

export const Route = createFileRoute("/")({
    component: HomeComponent,
});

function HomeComponent() {
    return (
        <main className="w-screen h-screen">
            <h3 className="text-2xl ml-4 mt-4">
                Home
            </h3>
        </main>
    );
}
