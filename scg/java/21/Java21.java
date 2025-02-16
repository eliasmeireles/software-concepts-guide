
import java.lang.ScopedValue;
import java.util.concurrent.StructuredTaskScope;

import static java.util.concurrent.Executors.newVirtualThreadPerTaskExecutor;
import static java.util.stream.IntStream.*;

/**
 * Exemplos práticos das principais features do Java 21.
 * Este arquivo contém demonstrações de:
 * 1. Virtual Threads
 * 2. Scoped Values
 * 3. Pattern Matching para switch
 * 4. Sequenced Collections
 * 5. String Templates (Preview)
 * 6. Generational ZGC (demonstração conceitual)
 */
public class Java21 {

    public static void main(String[] args) {
        System.out.println("=== Demonstração das Features do Java 21 ===");

        // Executa cada exemplo
        virtualThreadsExample();
        scopedValuesExample();
        patternMatchingSwitchExample();
        sequencedCollectionsExample();
        stringTemplatesExample();
        generationalZGCExample();
    }

    // 1. Virtual Threads
    private static void virtualThreadsExample() {
        System.out.println("\n--- Virtual Threads ---");

        try (var executor = newVirtualThreadPerTaskExecutor()) {
            range(0, 10).forEach(i ->
                    executor.submit(() -> {
                        System.out.println("Thread: " + i + " | Virtual Thread: " + Thread.currentThread().isVirtual());
                    })
            );
        }
    }

    // 2. Scoped Values
    private static void scopedValuesExample() {
        System.out.println("\n--- Scoped Values ---");

        final ScopedValue<String> USER_ID = ScopedValue.newInstance();

        // Define um valor escopo para USER_ID
        ScopedValue.where(USER_ID, "user123").run(() -> {
            System.out.println("Processing request for user: " + USER_ID.get());

            try (var scope = new StructuredTaskScope<>()) {
                var task1 = scope.fork(() -> fetchData("Data1", USER_ID.get()));
                var task2 = scope.fork(() -> fetchData("Data2", USER_ID.get()));

                scope.join(); // Espera todas as tarefas terminarem

                System.out.println("Task1 result: " + task1.get());
                System.out.println("Task2 result: " + task2.get());
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                System.err.println("Task interrupted");
            }
        });
    }

    private static String fetchData(String data, String userId) {
        return "Fetched " + data + " for user: " + userId;
    }

    // 3. Pattern Matching para switch
    private static void patternMatchingSwitchExample() {
        System.out.println("\n--- Pattern Matching para switch ---");

        Object obj = "Hello, Java 21!";

        String result = switch (obj) {
            case Integer i -> "Integer: " + i;
            case Long l -> "Long: " + l;
            case String s -> "String: " + s;
            default -> "Unknown type";
        };

        System.out.println(result); // Output: "String: Hello, Java 21!"
    }

    // 4. Sequenced Collections
    private static void sequencedCollectionsExample() {
        System.out.println("\n--- Sequenced Collections ---");

        java.util.SequencedCollection<String> sequencedCollection = new java.util.ArrayList<>();
        sequencedCollection.add("Primeiro");
        sequencedCollection.add("Segundo");
        sequencedCollection.add("Terceiro");

        System.out.println("First: " + sequencedCollection.getFirst()); // Output: "Primeiro"
        System.out.println("Last: " + sequencedCollection.getLast());  // Output: "Terceiro"
    }

    // 5. String Templates (Preview)
    private static void stringTemplatesExample() {
        System.out.println("\n--- String Templates (Preview) ---");

        String name = "Java 21";
        String message = STR."Hello, \{name}!";
        System.out.println(message); // Output: "Hello, Java 21!"
    }

    // 6. Generational ZGC (demonstração conceitual)
    private static void generationalZGCExample() {
        System.out.println("\n--- Generational ZGC ---");

        // Simula alocação de muitos objetos
        for (int i = 0; i < 1_000_000; i++) {
            String temp = new String("Object " + i);
        }
        System.out.println("Objects allocated and collected efficiently by Generational ZGC.");
    }
}
