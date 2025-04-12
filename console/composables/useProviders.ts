interface Provider {
    label: string
    name: string
    disabled: boolean
    icon: string
  }

  export function useProviders(currentProvider: string) {
    const providers = ref<Provider[]>([
        {
            label: 'Keycloak',
            name: 'keycloak',
            disabled: Boolean(currentProvider === 'keycloak'),
            icon: 'i-simple-icons-cncf',
          },
        ])
        return {
            providers,
          }
        }