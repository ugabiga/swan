import { NavigationMenu, NavigationMenuItem, NavigationMenuList } from '@/components/ui/navigation-menu'
import { Link, Outlet, createRootRoute, useMatchRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
  component: RootComponent,
})

function RootComponent() {
  const matchRoute = useMatchRoute()

  return (
    <>
      <div className="flex w-full border-b">
        <NavigationMenu className="w-full px-4 py-2">
          <NavigationMenuList className="w-full flex justify-between items-center">
            <NavigationMenuItem>
              <Link
                to="/"
                className={`pr-4 py-2 rounded-md transition-colors ${matchRoute({ to: '/' }) ? 'font-bold' : 'font-normal'
                  }`}
              >
                Home
              </Link>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <Link
                to="/about"
                className={`pr-4 py-2 rounded-md transition-colors ${matchRoute({ to: '/about' }) ? 'font-bold' : 'font-normal'
                  }`}
              >
                About
              </Link>
            </NavigationMenuItem>

          </NavigationMenuList>
        </NavigationMenu>

        <div className="flex-grow flex justify-end"></div>

        <div className='pr-1 py-1'>
        </div>
      </div>

      <Outlet />
      <TanStackRouterDevtools position="bottom-right" />
    </>
  )
}
