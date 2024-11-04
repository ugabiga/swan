import { createFileRoute } from "@tanstack/react-router";
import "../index.css";

export const Route = createFileRoute("/")({
    component: HomeComponent,
});

function HomeComponent() {
    return (
        <main className="p-4 space-y-4">
            <h3 className="text-2xl">Home</h3>
        </main>
    );
}
