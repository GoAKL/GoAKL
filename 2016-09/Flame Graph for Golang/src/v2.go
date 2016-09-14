func dfs(rbac *RBAC, id string, skipped map[string]struct{}, stack []string) error {
	if _, ok := skipped[id]; ok {
		return nil
	}
	for _, item := range stack {
		if item == id {
			return fmt.Errorf("Found circle: %s", stack)
		}
	}
	if len(rbac.parents[id]) == 0 {
		stack = make([]string, 0)
		skipped[id] = empty
		return nil
	}
	stack = append(stack, id)
	for pid := range rbac.parents[id] {
		if err := dfs(rbac, pid, skipped, stack); err != nil {
			return err
		}
	}
	return nil
}
