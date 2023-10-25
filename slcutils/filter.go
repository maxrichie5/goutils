package slcutils

// Remove removes items from a list of items using a custom function.
// If the remove function returns true, the item will be removed from
// the list of items and be put into the removed list.
// If the remove function returns false, the item will be put into the
// kept list.
//
// Remove does the opposite of [Keep]
func Remove[T any](items []T, shouldRemove func(T) bool) (removed, kept []T) {
	removed = make([]T, 0, len(items)/2)
	kept = make([]T, 0, len(items)/2)

	for _, item := range items {
		if shouldRemove(item) {
			removed = append(removed, item)
		} else {
			kept = append(kept, item)
		}
	}

	return
}

// Keep keeps items in a list of items using a custom function.
// If the keep function returns false, the item will be removed from
// the list of items and be put into the removed list.
// If the keep function returns true, the item will be put into the
// kept list.
//
// Keep does the opposite of [Remove].
func Keep[T any](items []T, shouldKeep func(T) bool) (kept, removed []T) {
	return Remove(items, shouldKeep)
}
