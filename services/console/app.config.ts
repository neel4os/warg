export default defineAppConfig({
    theme: 'dark',
    colors: {
        primary: '#3B82F6',
        secondary: '#4299E1',
        accent: '#D97706',
        error: '#EF4444',
        warning: '#F59E0B',
        info: '#6B7280',
        success: '#10B981'
    },
    icons: {
        iconfont: 'mdiSvg'
    },
    components: true,
    buildModules: [
        '@nuxt/typescript-build',
        '@nuxtjs/composition-api/module',
        '@nuxtjs/color-mode',
        '@nuxtjs/tailwindcss'
    ],
    modules: [
        '@nuxt/content',
        '@nuxt/image',
        '@nuxtjs/axios',
    ],
    content: {
        markdown: {
            prism: {
                theme: 'prism-themes/themes/prism-material-oceanic.css'
            }
        }
    },
    // image: {
    //     provider: 'static',
    //     dir: 'assets/images'
    // },
})