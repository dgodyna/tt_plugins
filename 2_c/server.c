#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

int main(int argc, char **argv) {
    void *handle;
    void (*func_echo)(const char *);

    handle = dlopen("libplugin.so", RTLD_LAZY);
    if (!handle) {
        fprintf(stderr, "Error: %s\n", dlerror());
        return EXIT_FAILURE;
    }

    *(void **) (&func_echo) = dlsym(handle, "echo");
    if (!func_echo) {
        fprintf(stderr, "Error: %s\n", dlerror());
        dlclose(handle);
        return EXIT_FAILURE;
    }

    func_echo(argv[1]);
    dlclose(handle);

    return EXIT_SUCCESS;
}